# Link: https://leetcode-cn.com/problems/last-person-to-fit-in-the-elevator


SELECT `a`.`person_name`
FROM (
         SELECT `person_name`, @`pre` := @`pre` + `weight` `weight`
         FROM `queue`,
              (SELECT @`pre` := 0) `tmp`
         ORDER BY `turn`
     ) `a`
WHERE `a`.`weight` <= 1000
ORDER BY `a`.`weight` DESC
LIMIT 1