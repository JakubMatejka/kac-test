CREATE TABLE "address_city" AS
SELECT "a"."name", "a"."surname", "c"."LatD", "c"."LonD"
FROM "addresses" "a"
LEFT JOIN "cities" "c" ON ("a"."city" = "c"."City")
WHERE "a"."state"={{state}}
