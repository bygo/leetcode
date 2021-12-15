# Link: https://leetcode-cn.com/problems/count-student-number-in-departments

SELECT `dept_name`, COUNT(`student_id`) `student_number`
FROM `department`            `d`
         LEFT JOIN `student` `s` ON `d`.`dept_id` = `s`.`dept_id`
GROUP BY `dept_name`
ORDER BY `student_number` DESC