# Title: Students Report By Geography
# Link: https://leetcode-cn.com/problems/students-report-by-geography

SELECT `America`, `Asia`, `Europe`
FROM (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) AS `r`, `name` AS `America`
      FROM `student`
      WHERE `continent` = 'America') `a`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) AS `r`, `name` AS `Asia`
                    FROM `student`
                    WHERE `continent` = 'Asia') `b` ON `a`.`r` = `b`.`r`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) AS `r`, `name` AS `Europe`
                    FROM `student`
                    WHERE `continent` = 'Europe') `c` ON `a`.`r` = `c`.`r`