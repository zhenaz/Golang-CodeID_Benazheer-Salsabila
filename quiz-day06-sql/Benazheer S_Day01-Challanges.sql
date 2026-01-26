-- Benazheer Salsabila

-- 2.1 1. Tampilkan data category & total product-nya.
select c.category_id, category_name, count(product_id) as total_product
from oe.categories c
join oe.products p on c.category_id = p.category_id 
group by c.category_id, category_name
order by category_name

--  2.2 Tampilkan data supplier & total product-nya.
select s.supplier_id, company_name, count(product_id) as total_product
from oe.suppliers s
join oe.products p on s.supplier_id = p.supplier_id 
group by s.supplier_id, company_name
order by total_product desc

-- 2.3. Tampilkan data supplier, total product dan harga rata-rata tiap product (gunakaan to_char() untuk menampilkan format avg_unit_price).
select s.supplier_id, company_name, count(product_id) as total_product, to_char(round(avg(unit_price)::numeric, 2), 'FM999999.00') as avg_unit_price
from oe.suppliers s
join oe.products p on s.supplier_id = p.supplier_id 
group by s.supplier_id, company_name
order by total_product desc

-- 2.4. Tampilkan data product yang harus pesan lagi ke supplier sesuai reorder-level nya.
select product_id, product_name, p.supplier_id, company_name, unit_price, units_in_stock, units_on_order, reorder_level
from oe.products p
join oe.suppliers s on s.supplier_id = p.supplier_id 
where units_in_stock <= reorder_level 
order by product_name

-- 2.5. Tampilkan data customer dan total order-nya
select c.customer_id, company_name, count(o.customer_id) as total_order
from oe.customers c
join oe.orders o on c.customer_id = o.customer_id
group by c.customer_id, company_name
order by c.customer_id

-- 2.6. Tampilkan data order yang melebihi delivery time lebih dari 7 hari.
select order_id, customer_id, order_date, required_date, shipped_date, (shipped_date - order_date) as delivery_time
from oe.orders
where shipped_date - order_date > 7
order by order_id

-- 2.7. Tampilkan total product yang sudah di order dan urut berdasarkan total_quantity terbesar

select p.product_id, product_name, sum(quantity) as total_qty
from oe.products p
join oe.order_details od on p.product_id = od.product_id
group by p.product_id, product_name
order by total_qty desc

--2.8 Tampilkan total product yang sudah di order berdasarkan category
select c.category_id, category_name, sum(od.quantity) as total_qty_ordered
from oe.categories c
join oe.products p on p.category_id = c.category_id
join oe.order_details od on p.product_id = od.product_id
group by c.category_id, category_name
order by total_qty_ordered desc


--2.9 Mengacu ke soal no 8, tampilkan data category yang memiliki min & max total_qty_ordered
with category_qty as
	(select c.category_id, category_name, sum(od.quantity) as total_qty_ordered
	from oe.categories c
	join oe.products p on p.category_id = c.category_id
	join oe.order_details od on p.product_id = od.product_id
	group by c.category_id, category_name
	)
	
select category_id, category_name, total_qty_ordered
from category_qty
where total_qty_ordered = (select min(total_qty_ordered) from category_qty) or
	total_qty_ordered = (select max(total_qty_ordered) from category_qty)
order by total_qty_ordered desc

--2.10. Tampilkan data shipper dan total qty product yang dikirim

select shipper_id, company_name, p.product_id, product_name, sum(od.quantity) as total_qty_ordered
from oe.orders o
join oe.shippers s on o.ship_via = s.shipper_id
join oe.order_details od on od.order_id = o.order_id
join oe.products p on p.product_id = od.product_id
group by shipper_id, company_name, p.product_id, product_name
order by shipper_id desc, product_name

--2.11 Tampilkan data shipper dengan product yang paling sering dikirim dan paling minim dikirim.

with shipper_qty as
	(select shipper_id, company_name, p.product_id, product_name, sum(od.quantity) as total_qty_ordered
	from oe.orders o
	join oe.shippers s on o.ship_via = s.shipper_id
	join oe.order_details od on od.order_id = o.order_id
	join oe.products p on p.product_id = od.product_id
	group by shipper_id, company_name, p.product_id, product_name
	order by shipper_id desc, product_name
	)
	
select shipper_id, company_name, product_id, product_name, total_qty_ordered
from shipper_qty
where total_qty_ordered = (select min(total_qty_ordered) from shipper_qty where shipper_id =1)  or
	total_qty_ordered = (select max(total_qty_ordered) from shipper_qty where shipper_id =1) or
	total_qty_ordered = (select min(total_qty_ordered) from shipper_qty where shipper_id =2)  or
	total_qty_ordered = (select max(total_qty_ordered) from shipper_qty where shipper_id =2)  or
	total_qty_ordered = (select min(total_qty_ordered) from shipper_qty where shipper_id =3)  or
	total_qty_ordered = (select max(total_qty_ordered) from shipper_qty where shipper_id =3)
order by shipper_id, total_qty_ordered asc