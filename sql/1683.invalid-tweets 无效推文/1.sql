# Link: https://leetcode-cn.com/problems/invalid-tweets

SELECT `tweet_id`
FROM `tweets`
WHERE char_length(`content`) > 15