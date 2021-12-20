# Link: https://leetcode-cn.com/problems/create-a-session-bar-chart


SELECT `a`.`bin`, COUNT(`b`.`bin`) `total`
FROM (
         SELECT '[0-5>' `bin` UNION SELECT '[5-10>' `bin` UNION SELECT '[10-15>' `bin` UNION SELECT '15 or more' `bin`
     ) `a`
         LEFT JOIN
     (
         SELECT CASE
                    WHEN `duration` < 300 THEN '[0-5>'
                    WHEN `duration` >= 300 AND `duration` < 600 THEN '[5-10>'
                    WHEN `duration` >= 600 AND `duration` < 900 THEN '[10-15>'
                    ELSE '15 or more'
                    END `bin`
         FROM `sessions`
     ) `b`
     ON `a`.`bin` = `b`.`bin`
GROUP BY `a`.`bin`