# Title: Investments in 2016
# Link: https://leetcode-cn.com/problems/investments-in-2016

SELECT ROUND(SUM(`tiv_2016`), 2) AS `tiv_2016`
FROM (SELECT *,
             COUNT(*) OVER ( PARTITION BY `tiv_2015`)  AS `y`,
             COUNT(*) OVER ( PARTITION BY `lat`,`lon`) AS `p`
      FROM `insurance`) `t`
WHERE `y` > 1 && `p` = 1