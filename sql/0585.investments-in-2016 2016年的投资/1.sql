# Link: https://leetcode.cn/problems/investments-in-2016

SELECT ROUND(SUM(`tiv_2016`), 2) `tiv_2016`
FROM (SELECT *,
             COUNT(*) OVER ( PARTITION BY `tiv_2015`)  `y`,
             COUNT(*) OVER ( PARTITION BY `lat`,`lon`) `p`
      FROM `insurance`) `t`
WHERE `y` > 1 && `p` = 1