# Link: https://leetcode.cn/problems/rank-scores

SELECT `score`,
       @`rank` := @`rank` + (@`pre` != (@`pre` := `score`)) `rank`
FROM `scores`,
     (SELECT @`pre` := -1, @`rank` := 0) `tmp`
ORDER BY `score` DESC;

#

SELECT `score`,
       DENSE_RANK() OVER (ORDER BY `score` DESC) `rank`
FROM `scores`