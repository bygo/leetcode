# Link: https://leetcode-cn.com/problems/second-degree-follower

SELECT `followee`                 AS `follower`,
       count(DISTINCT `follower`) AS `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`



SELECT `followee`                 AS `follower`,
       count(DISTINCT `follower`) AS `num`
FROM `follow`
WHERE `followee` IN (SELECT `follower` FROM `follow`)
GROUP BY `followee`
ORDER BY `followee`