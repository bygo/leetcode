# Title: Exchange Seats
# Link: https://leetcode-cn.com/problems/exchange-seats

SELECT row_number() OVER (ORDER BY (`id` + 1 - 2 * power(0, `id` % 2))) AS `id`,
       `student`
FROM `seat`



SELECT IF(`id` % 2 = 0, `id` - 1, `id` + 1) AS `id`
    `student`
FROM `ORDER` BY id ASC;
