$(document).ready(function () {

      $("#tableUser").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableRoles").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableGender").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableBanks").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableDocument").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableEvidence").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableRT").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableSampah").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableJadwal").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableKepengurusan").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

          $("#tableUserTransaction").DataTable({
          pagingType: "full_numbers", // Penomoran halaman lengkap
          lengthMenu: [5, 10, 25, 50], // Pilihan jumlah data per halaman
          language: {
            url: "https://cdn.datatables.net/plug-ins/1.13.7/i18n/id.json", // Translate ke Bahasa Indonesia
          },
          order: [[0, "desc"]], // Urutkan berdasarkan ID paling baru
        });

    const $card = $('#card_project');

    const $icon = $('.card-header').find('i');
    $(document).on('click', '.btn-tool', function () {

        // toggle collapsed state
        const isCollapsed = $card.hasClass('collapsed-card');

        if (isCollapsed) {
            // buka card
            $card.removeClass('collapsed-card');
            $card.find('.card-body').slideDown(200);
            $icon.removeClass('fa-plus').addClass('fa-minus');
            $card.find('.btn-cancel').removeClass('hidden');
        } else {
            // tutup card
            $card.addClass('collapsed-card');
            $card.find('.card-body').slideUp(200);
            $icon.removeClass('fa-minus').addClass('fa-plus');
            $card.find('.btn-cancel').addClass('hidden');
        }

        $('#formUser input[name="_method"]').val("POST");
        $('#formUser input[name="id"]').val("");
        $('#formUser input[name="fullName"]').val("");
        $('#formUser input[name="email"]').val("");
    });

    $(document).on('click', '.btn-cancel', function () {

        const isCollapsed = $card.hasClass('collapsed-card');

        // tutup card
        $card.addClass('collapsed-card');
        $card.find('.card-body').slideUp(200);
        $icon.removeClass('fa-minus').addClass('fa-plus');
        $card.find('.btn-cancel').addClass('hidden');
        $card.find('.btn-simpan').text('Simpan Buku')

    });
    $('#formUser').on('submit', function (e) {
        e.preventDefault();

        const method = $('#formUser input[name="_method"]').val();
        const id = $('#formUser input[name="id"]').val();


        const url = method === 'POST'

        try {
            const post = () => {
                const formData = new FormData(this);


                $.ajax({
                    url: url,
                    method: 'POST', // tetap POST
                    data: formData,
                    contentType: false,
                    processData: false,
                    headers: {
                        'Accept': 'application/json',
                        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                    },
                    success: function (result) {
                        Swal.fire({
                            icon: 'success',
                            title: 'Success!',
                            text: result.message,
                            timer: 1500,
                            showConfirmButton: false
                        }).then(() => location.reload());
                    },
                    error: function (xhr) {
                        if (xhr.status === 422) {

                            Swal.fire('Gagal!', 'Silakan periksa kembali inputan Anda.', 'error');
                        } else {
                            Swal.fire('Error', xhr.responseJSON?.message || 'Server error', 'error');
                        }
                    },
                    complete: function () {
                        Swal.fire({
                            icon: 'success',
                            title: 'Success!',
                            text: result.message,
                            timer: 1500,
                            showConfirmButton: false
                        }).then(() => location.reload());
                    }
                });
            };


            method === 'POST' ?
                Swal.fire({
                    title: 'Are you sure?',
                    text: "Do you want to add this people?",
                    icon: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#d33',
                    cancelButtonColor: '#3085d6',
                    confirmButtonText: 'Yes, add it!'
                }).then(() => post()) : Swal.fire({
                    title: 'Are you sure?',
                    text: "Do you want to update this people?",
                    icon: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#d33',
                    cancelButtonColor: '#3085d6',
                    confirmButtonText: 'Yes, update it!'
                }).then((result) => {
                    Swal.fire({
                        icon: 'success',
                        title: 'Update!',
                        text: result.message,
                        timer: 1500,
                        showConfirmButton: false
                    }).then(() => post());


                })
        } catch (error) {
            if (typeof show_error === 'function') show_error({
                html: error
            });
            else alert(error);
        }
    });


    // Edit button - delegated (works with DataTables)
    $(document).on('click', '#tableUser .btn-edit', function () {
        const data = $(this).data();
        const isCollapsed = $('#card_project').hasClass('collapsed-card');
        // open card
        $('#card_project').removeClass('collapsed-card');


        if (isCollapsed) {
            $card.removeClass('collapsed-card');
            $card.find('.card-body').slideDown(200);
            $card.find('.btn-cancel').removeClass('hidden');

            $card.find('.btn-simpan').text('Update Nasabah')
        } else {
            $card.addClass('collapsed-card');
            $card.find('.card-body').slideUp(200);
            $card.find('.btn-cancel').addClass('hidden');
        }
        $('#card_project .btn-cancel').removeClass('hidden');
        // isi form
        $('#formUser input[name=id]').val(data.id);
        $('#formUser input[name="_method"]').val("PUT"); // Ubah method jadi PUT
        $('#formUser input[name="fullname"]').val(data.fullname);
        $('#formUser input[name="email"]').val(data.email)
        $('#formUser input[name="address"]').val(data.address)
        $('#formUser input[name="telephone"]').val(data.telephone)
        $('#formUser select[name="status"]').val(data.status)
    });

    // Cancel button
    $(document).on('click', '#card_project .btn-cancel', function () {
        $('#formUser input[name="id"]').val("");
        $('#formUser input[name="_method"]').val("POST");
        $('#card_project .btn-cancel').addClass('hidden');
        // close form
        $('#card_project').addClass('collapsed-card');
        $('#card_project .card-body').hide();
        $('#card_project .btn-tool>i').removeClass('fa-minus').addClass('fa-plus');
    });

    $(document).on('click', '.btn-delete', function () {
        const id = $(this).data('id');
        const token = $('meta[name="csrf-token"]').attr('content');

        if (!id) {
            alert('ID tidak ditemukan!');
            return;
        }

        const post = () => {
            $.ajax({
                url: ``,
                type: 'POST',
                data: {
                    _method: 'DELETE',
                    _token: token
                },
                headers: {
                    'Accept': 'application/json',
                    'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
                },
                success: function (result) {
                    if (result.code === 200) {
                        Swal.fire({
                            icon: 'success',
                            title: 'Deleted!',
                            text: result.message,
                            timer: 1500,
                            showConfirmButton: false
                        }).then(() => location.reload());
                    } else {
                        Swal.fire('Error', result.message ?? 'Gagal menghapus data',
                            'error');
                    }
                },
                error: function (xhr) {
                    Swal.fire('Error', xhr.responseJSON?.message ||
                        'Terjadi kesalahan server', 'error');
                }
            });
        };

        Swal.fire({
            title: 'Are you sure?',
            text: "Do you want to delete this book?",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#d33',
            cancelButtonColor: '#3085d6',
            confirmButtonText: 'Yes, delete it!'
        }).then((result) => {
            if (result.isConfirmed) post();
        });
    });
})