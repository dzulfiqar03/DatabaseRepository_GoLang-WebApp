package HomeController

import (
	dbgivent "database-repository/services/db-givent"
	"database-repository/services/db-libraryis"
	"database-repository/services/db-netabot"
	dbseeu "database-repository/services/db-seeu"
	"fmt"
	"net/http"
	"sort" // Diperlukan untuk meranking data
	"strconv"

	"github.com/gin-gonic/gin"
)

// LeaderboardEntry untuk menampung data peringkat
type LeaderboardEntry struct {
	Name  string
	Label string
	Value string
	Score float64
}

// SummaryRow untuk tabel rangkuman antar sistem
type SummaryRow struct {
	SystemName string
	MainMetric string
	Volume     string
	Status     string
	Color      string
}

// TableSection penampung utama data ke template
type TableSection struct {
	Title        string
	TotalBooks   int
	TotalFine    float64
	TotalNeta    int
	TotalUMKM    int
	TotalCust    int
	TitleTable   string
	NetaLabels   []string
	NetaSold     []int
	NetaRating   []float64
	SeeULabels   []string
	SeeUData     []int
	SampahLabels []string
	SampahData   []int
	SummaryTable []SummaryRow
	TopWarga     []LeaderboardEntry
	TopRT        []LeaderboardEntry
	RecentAct    []string
}


func Index(c *gin.Context) {
	// 1. Ambil Data Real dari Database
	dataBook, _, _ := dblibraryis.GetBooks()
	dataFine, _, _ := dblibraryis.GetTransaction()
	dataProdNeta, _, _ := dbnetabot.GetProduct()
	dataAllUmkm, _, _ := dbseeu.GetAllUmkm()
	dataCust, _, _ := dbgivent.GetCustomer()

	// 2. Library Analytics (Total Denda)
	var totalFine float64
	for _, f := range dataFine {
		totalFine += f.Fine_Amount
	}

		var dataSidebar = []string{
		"db-sibanksa",
		"db-givent",
		"db-seeU",
		"db-netabot",
		"db-libraryis",
	}

	// 3. SeeU Analytics (Dinamis Kategori)
	catMap := make(map[string]int)
	for _, u := range dataAllUmkm {
		name := "Lainnya"
		if u.Category.Name != "" { name = u.Category.Name }
		catMap[name]++
	}
	var seeuLabels []string
	var seeuData []int
	for k, v := range catMap {
		seeuLabels, seeuData = append(seeuLabels, k), append(seeuData, v)
	}

	// 4. Netabot Analytics (Top 5 Produk)
	var netaLabels []string
	var netaSold []int
	var netaRating []float64
	for i, p := range dataProdNeta {
		if i >= 5 { break }
		netaLabels = append(netaLabels, p.Name)
		s, _ := strconv.Atoi(p.Sold)
		netaSold = append(netaSold, s)
		r, _ := strconv.ParseFloat(p.Rating, 64)
		netaRating = append(netaRating, r)
	}

	// 5. GENERATOR DATA DUMMY SIBANKSA (50 Transaksi)
	// Kita simulasikan 50 setoran dengan RT dan Kategori Sampah acak
	type SetoranItem struct {
		Kategori string
		Berat    int // dalam gram
	}
	type SetoranDummy struct {
		NamaWarga string
		RT        string
		Total     int
		Items     []SetoranItem
	}

	var dummyTransactions []SetoranDummy
	rtOptions := []string{"RT A", "RT B", "RT C"}
	kategoriOptions := []string{"Plastik", "Kertas", "Logam", "Kaca", "Organik"}

	for i := 1; i <= 50; i++ {
		rtTarget := rtOptions[i%3]
		dummyTransactions = append(dummyTransactions, SetoranDummy{
			NamaWarga: fmt.Sprintf("Warga %02d", i),
			RT:        rtTarget,
			Total:     (i * 1500) % 75000,
			Items:     []SetoranItem{{Kategori: kategoriOptions[i%5], Berat: (i * 250) % 5000}},
		})
	}

	// 6. PROSES LEADERBOARD TOP 5 (Khusus RT A)
	wargaRTAMap := make(map[string]int)
	sampahRTAMap := make(map[string]int)

	for _, t := range dummyTransactions {
		if t.RT == "RT A" {
			wargaRTAMap[t.NamaWarga] += t.Total
			for _, item := range t.Items {
				sampahRTAMap[item.Kategori] += item.Berat
			}
		}
	}

	// Sorting Warga RT A
	var top5WargaRTA []LeaderboardEntry
	for name, total := range wargaRTAMap {
		top5WargaRTA = append(top5WargaRTA, LeaderboardEntry{
			Name:  name,
			Value: fmt.Sprintf("Rp %d", total),
			Score: float64(total),
		})
	}
	sort.Slice(top5WargaRTA, func(i, j int) bool {
		return top5WargaRTA[i].Score > top5WargaRTA[j].Score
	})
	if len(top5WargaRTA) > 5 { top5WargaRTA = top5WargaRTA[:5] }

	// Normalisasi Progress Bar Score (Juara 1 = 100%)
	if len(top5WargaRTA) > 0 {
		max := top5WargaRTA[0].Score
		for i := range top5WargaRTA {
			top5WargaRTA[i].Score = (top5WargaRTA[i].Score / max) * 100
		}
	}

	// 7. Komposisi Sampah RT A untuk Chart
	var sLabels []string
	var sData []int
	for k, v := range sampahRTAMap {
		sLabels, sData = append(sLabels, k), append(sData, v/1000) // Konversi ke Kg
	}

	// 8. Susun Rangkuman Akhir
	summary := []SummaryRow{
		{"Library IS", "Buku & Denda", strconv.Itoa(len(dataBook)), "Online", "blue"},
		{"Netabot", "E-Commerce", strconv.Itoa(len(dataProdNeta)), "Aktif", "indigo"},
		{"SeeU", "Partner UMKM", strconv.Itoa(len(dataAllUmkm)), "Stabil", "emerald"},
		{"Sibanksa", "Bank Sampah", "50 Transaksi", "Aktif", "amber"},
	}

	sections := []TableSection{
		{
			Title:        "Executive Dashboard - RT A Focus",
			TotalBooks:   len(dataBook),
			TotalFine:    totalFine,
			TotalNeta:    len(dataProdNeta),
			TotalUMKM:    len(dataAllUmkm),
			TotalCust:    len(dataCust),
			NetaLabels:   netaLabels,
			NetaSold:     netaSold,
			NetaRating:   netaRating,
			SeeULabels:   seeuLabels,
			SeeUData:     seeuData,
			SampahLabels: sLabels,
			SampahData:   sData,
			SummaryTable: summary,
			TopWarga:     top5WargaRTA,
			RecentAct:    []string{"RT A menyetor 12kg Kertas", "Buku baru ditambahkan", "Transaksi Netabot sukses"},
		},
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"page": "Dashboard", 
		"title":          "Halaman Dashboard",
		"arrDataSidebar": dataSidebar,
		"sections": sections})
}