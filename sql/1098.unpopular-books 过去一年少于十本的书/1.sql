# Link: https://leetcode-cn.com/problems/unpopular-books


SELECT `b`.`book_id`, `b`.`name`
FROM `books` `b`
         LEFT JOIN `orders` `o` ON `b`.`book_id` = `o`.`book_id`
    AND '2018-06-23' <= `o`.`dispatch_date`
WHERE `b`.`available_from` < '2019-05-23'
GROUP BY `b`.`book_id`
HAVING ifnull(sum(`o`.`quantity`), 0) < 10