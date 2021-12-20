# Link: https://leetcode-cn.com/problems/find-interview-candidates

WITH `a` AS
         (SELECT DISTINCT `temp1`.`name`, `mail`
          FROM (
                   SELECT `contest_id`,
                          `u1`.`user_id`,
                          `u1`.`name`,
                          `mail`,
                          `contest_id` - ROW_NUMBER() OVER (PARTITION BY `u1`.`user_id` ORDER BY `contest_id`) `diffs`
                   FROM `users`             `u1`
                            JOIN `contests` `c`
                                 ON `u1`.`user_id` IN (`c`.`gold_medal`, `c`.`silver_medal`, `c`.`bronze_medal`)
                   ORDER BY `u1`.`user_id`, `contest_id`) `temp1`
          GROUP BY `temp1`.`user_id`, `diffs`
          HAVING COUNT(`contest_id`) >= 3
          ORDER BY `temp1`.`user_id`)
SELECT *
FROM `a`
UNION
SELECT `u2`.`name`, `mail`
FROM `users`             `u2`
         JOIN `contests` `c1`
              ON `u2`.`user_id` = `c1`.`gold_medal`
GROUP BY `c1`.`gold_medal`
HAVING COUNT(*) >= 3

