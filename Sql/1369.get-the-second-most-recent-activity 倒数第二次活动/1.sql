# Link: https://leetcode.cn/problems/get-the-second-most-recent-activity

SELECT `username`, `activity`, `startdate`, `enddate`
FROM (SELECT *,
             RANK() OVER (PARTITION BY `username` ORDER BY `startdate` DESC) `r`,
             COUNT(*) OVER (PARTITION BY `username`)                         `count`
      FROM `useractivity`) `t`
WHERE `r` = 2
   OR `count` = 1;

