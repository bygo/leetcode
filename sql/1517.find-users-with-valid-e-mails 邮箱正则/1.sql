# Link: https://leetcode.cn/problems/find-users-with-valid-e-mails

SELECT *
FROM `users`
WHERE `mail` REGEXP '^[a-zA-Z][a-zA-Z0-9\_\.\-]*@leetcode\.com$'