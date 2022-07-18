# Link: https://leetcode.cn/problems/exchange-seats

SELECT ROW_NUMBER() OVER (ORDER BY (`id` + 1 - 2 * POWER(0, `id` % 2))) `id`,
       `student`
FROM `seat`



SELECT IF(`id` % 2 = 0, `id` - 1, `id` + 1) `id`
    `student`
FROM `order` BY id ASC;
