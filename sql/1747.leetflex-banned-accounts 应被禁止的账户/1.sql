# Link: https://leetcode.cn/problems/leetflex-banned-accounts

SELECT DISTINCT `l1`.`account_id`
FROM `loginfo`          `l1`
         JOIN `loginfo` `l2`
              ON `l1`.`account_id` = `l2`.`account_id`
                  AND `l1`.`ip_address` <> `l2`.`ip_address`
                  AND ((`l1`.`login` BETWEEN `l2`.`login` AND `l2`.`logout`)
                      OR (`l1`.`logout` BETWEEN `l2`.`login` AND `l2`.`logout`))
