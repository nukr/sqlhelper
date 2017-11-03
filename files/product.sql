-- name: createProduct

INSERT INTO products(
  price,
  title
)
VALUES (
  $1,
  $2
)

-- name: deleteProduct

DELETE FROM products
WHERE id = $1
