# Link: https://leetcode-cn.com/problems/delete-duplicate-emails

DELETE `p2`
FROM `Person` `p1`
         JOIN `Person` `p2`
WHERE `p1`.`Email` = `p2`.`Email`
  AND `p1`.`Id` < `p2`.`Id`