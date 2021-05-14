# Link: https://leetcode-cn.com/problems/bank-account-summary-ii


SELECT `name`,
       sum(`amount`) `balance`
FROM `users`,
     `transactions`
WHERE `users`.`account` = `transactions`.`account`
GROUP BY `name`
HAVING sum(`amount`) > 10000;


