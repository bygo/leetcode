# Link: https://leetcode-cn.com/problems/invalid-tweets

SELECT `tweet_id`
FROM `tweets`
WHERE CHAR_LENGTH(`content`) > 15