# Link: https://leetcode-cn.com/problems/bank-account-summary-ii


SELECT `name`,
       SUM(`amount`) `balance`
FROM `users`,
     `transactions`
WHERE `users`.`account` = `transactions`.`account`
GROUP BY `name`
HAVING SUM(`amount`) > 10000;


