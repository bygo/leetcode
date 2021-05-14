# Link: https://leetcode-cn.com/problems/count-apples-and-oranges


SELECT sum(`b`.`apple_count` + ifnull(`c`.`apple_count`, 0))   `apple_count`,
       sum(`b`.`orange_count` + ifnull(`c`.`orange_count`, 0)) `orange_count`
FROM `boxes`                `b`
    LEFT JOIN `chests` `c` ON `b`.`chest_id` = `c`.`chest_id`