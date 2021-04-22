# Link: https://leetcode-cn.com/problems/rank-scores

SELECT `score`,
       @`rank` := @`rank` + (@`pre` != (@`pre` := `score`)) `Rank`
FROM `scores`,
     (SELECT @`pre` := -1, @`rank` := 0) `tmp`
ORDER BY `score` DESC;

#

SELECT `score`,
       DENSE_RANK() OVER (ORDER BY `Score` DESC) `Rank`
FROM `scores`