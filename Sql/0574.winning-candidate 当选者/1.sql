# Link: https://leetcode.cn/problems/winning-candidate

SELECT `name`
FROM (
         SELECT `candidateid`
         FROM `vote`
         GROUP BY `candidateid`
         ORDER BY COUNT(`candidateid`) DESC
         LIMIT 1) `t`
         JOIN `candidate` ON `candidate`.`id` = `candidateid`