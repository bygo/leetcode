# Link: https://leetcode-cn.com/problems/students-report-by-geography

SELECT `america`, `asia`, `europe`
FROM (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `america`
      FROM `student`
      WHERE `continent` = 'America')              `a`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `asia`
                    FROM `student`
                    WHERE `continent` = 'Asia')   `b` ON `a`.`r` = `b`.`r`
         LEFT JOIN (SELECT `name`, ROW_NUMBER() OVER (ORDER BY `name`) `r`, `name` `europe`
                    FROM `student`
                    WHERE `continent` = 'Europe') `c` ON `a`.`r` = `c`.`r`