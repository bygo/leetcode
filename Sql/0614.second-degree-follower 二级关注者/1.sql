# Link: https://leetcode.cn/problems/second-degree-follower

SELECT `followee`                 `follower`,
       COUNT(DISTINCT `follower`) `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`



SELECT `followee`                 `follower`,
       COUNT(DISTINCT `follower`) `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`