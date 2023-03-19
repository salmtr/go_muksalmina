-- INSERT
-- Insert 5 operators pada table operators.

INSERT INTO operator(name) VALUES
('TOKO SAKTI'),
('TOKO MANDRAGUNA'),
('TOKO SEHAT'),
('TOKO KUAT'),
('TOKO SUKSES');

-- Insert 3 product type.

INSERT INTO product_types(name) VALUES
('PAKAIAN'),
('BAHAN BANGUNAN'),
('ELEKTRONIK');

-- Insert 2 product dengan product type id = 1, dan operators id = 3.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(1, 3, 'M001', 'semen', 1),
(1, 3, 'M002', 'besi', 1);

-- Insert 3 product dengan product type id = 2, dan operators id = 1.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(2, 1, 'N001', 'paku', 1),
(2, 1, 'N002', 'keramik', 1),
(2, 1, 'N003', 'genteng', 1);

-- Insert 3 product dengan product type id = 3, dan operators id = 4.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(3, 4, 'S004', 'TELEVISI', 1),
(3, 4, 'S005', 'KULKAS', 1),
(3, 4, 'S006', 'KIPAS ANGIN', 1);

-- Insert product description pada setiap product.

INSERT INTO product_descriptions(description) VALUES
('semen lima roda '),
('besi kuat dan tahan lama'),
('paku rumah bukan paku alam'),
('keramik untuk lantai'),
('genteng anti bocor'),
('televisi menonton saluran anda'),
('kulkas cool'),
('kipas angin membuat anda tidur nyenyak');

-- Insert 3 payment methods.

INSERT INTO payment_methods(name, status) VALUES
('Bsi', 1),
('Btn', 2),
('Bank Aceh', 3);

-- Insert 5 user pada tabel user.

INSERT INTO users(name, status, dob, gender) VALUES
('balmon', 1, '1980-02-07', 'L'),
('argus', 2, '1990-02-04', 'L'),
('odet', 3, '1997-06-16', 'p'),
('nana', 4, '2004-09-26', 'P'),
('gusion', 5, '2001-08-05', 'L');

-- Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

INSERT INTO transaksi(user_id, payments_method_id, status, total_qty, total_price) VALUES
(1, 1, 'Cash', 2, 100000),
(2, 2, 'Cash', 3, 200000),
(3, 3, 'Cash', 4, 89000);

-- Insert 3 product di masing-masing transaksi.

INSERT INTO product(product_type_id, operator_id, code, name, status) VALUES
(1, 1, 'S004', 'TELEVISI', 1),
(2, 2, 'S005', 'KULKAS', 1),
(3, 3, 'S006', 'KIPAS ANGIN', 1);

-- SELECT
-- Tampilkan nama user / pelanggan dengan gender Laki-laki / M.

SELECT name FROM users WHERE gender = 'L' or gender = 'M';

-- Tampilkan product dengan id = 3.

SELECT * FROM product WHERE id = 3;

-- Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.

SELECT * FROM users WHERE create_at BETWEEN DATE_SUB(NOW(), INTERVAL 7 DAY) AND NOW() AND name LIKE '%a%';

-- Hitung jumlah user / pelanggan dengan status gender Perempuan.

SELECT COUNT(*) FROM users WHERE gender = 'P';

-- Tampilkan data pelanggan dengan urutan sesuai nama abjad

SELECT * FROM users ORDER BY name ASC;

-- Tampilkan 5 data pada data product

SELECT * FROM product LIMIT 5;

-- UPDATE
-- Ubah data product id 1 dengan nama ‘product dummy’.

UPDATE product SET name = 'product dummy' WHERE id = 1;

-- Update qty = 3 pada transaction detail dengan product id = 1.

UPDATE transaksi_details SET qty = 3 WHERE product_id = 1;

-- DELETE
-- Delete data pada tabel product dengan id = 1.

DELETE FROM product WHERE id = 1;

-- Delete pada pada tabel product dengan product type id 1.

DELETE FROM product WHERE product_type_id = 1;


-- Join, Union, Sub query, Function


-- Gabungkan data transaksi dari user id 1 dan user id 2. --
SELECT * FROM transaction WHERE user_id = 1 OR user_id = 2;
 
-- Tampilkan jumlah harga transaksi user id 1.
SELECT SUM(total_price) FROM transaction WHERE user_id = 1;

-- Tampilkan semua field table product dan field name table product type yang saling berhubungan. -- 
select count(*) from transaction t
join transaction_details td on t.id = td.transaction_id
join product p on p.id = td.product_id
where p.product_type_id = 2;

-- Tampilkan semua field table transaction, field name table product dan field name table user. --
select t.*, p.nama as product, u.nama as users from transaction_details td
join transaction t on t.id = td.transaction_id
join users u on u.id = t.users_id
join product p on p.id = td.product_id
where p.deleted_at is null


 -- Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud. --
 delimiter $$
 create trigger delete_transaction_details
 before delete on transaction for each row
 begin
 declare trans_id int;
 set trans_id = old.id;
 delete from transaction_details where transaction_id = trans_id;
 end $$
 
 delete from transaction where id = 4;
 
 -- Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus. --
 delimiter $$
 create trigger update_transaction_id
 after delete on transaction_details for each row
 begin
 declare trans_id int;
 set trans_id = old.transaction_id;
 update transaction set total_quanty = (select sum(quanty) from transaction_details where transaction_id = trans_id)
 where id = trans_id;
 end $$
 
 select * from transaction;
 select * from transaction_details where transaction_id = 10;
 delete from transaction_details where transaction_id = 10 AND product_id = 5;
 
 -- Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query. -- 
 select * from product where id not in(
 select product_id from transaction_details
 );
 
select * from product where id in (
select p.id from product p
left outer join transaction_details td on p.id = td.product_id
where td.transaction_id is null
);
