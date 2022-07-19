# Link: https://leetcode.cn/problems/npv-queries

SELECT `q`.`id`,
       `q`.`year`,
       IFNULL(`npv`, 0) `npv`
FROM `queries` `q`
         LEFT JOIN
     `npv`     `n` USING (`id`, `year`)
