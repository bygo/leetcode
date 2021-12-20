# Link: https://leetcode-cn.com/problems/count-apples-and-oranges


SELECT SUM(`b`.`apple_count` + IFNULL(`c`.`apple_count`, 0))   `apple_count`,
       SUM(`b`.`orange_count` + IFNULL(`c`.`orange_count`, 0)) `orange_count`
FROM `boxes`                `b`
         LEFT JOIN `chests` `c` ON `b`.`chest_id` = `c`.`chest_id`