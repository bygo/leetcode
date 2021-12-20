# Link: https://leetcode-cn.com/problems/delete-duplicate-emails

DELETE `p2`
FROM `person`          `p1`
         JOIN `person` `p2`
WHERE `p1`.`email` = `p2`.`email`
  AND `p1`.`id` < `p2`.`id`