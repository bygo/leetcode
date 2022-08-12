# Link: https://leetcode.cn/problems/human-traffic-of-stadium

WITH `countt` AS (SELECT `id`,
                         COUNT(*) OVER (PARTITION BY `rn` ORDER BY `rn` ) `counter`
                  FROM (SELECT `id`,
                               `id` - ROW_NUMBER() OVER (ORDER BY `id`) `rn`
                        FROM `stadium`
                        WHERE `people` >= 100) `rowt`)
SELECT `s`.*
FROM `stadium` `s`
         JOIN `countt` ON
    `s`.`id` = `countt`.`id`
WHERE `countt`.`counter` > 2
ORDER BY `s`.`visit_date`;