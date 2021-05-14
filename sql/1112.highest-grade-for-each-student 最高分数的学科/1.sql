# Link: https://leetcode-cn.com/problems/highest-grade-for-each-student

SELECT `student_id`, `course_id`, `grade`
FROM (SELECT *, row_number() OVER (PARTITION BY `student_id` ORDER BY `grade` DESC,`course_id`) AS `r`
      FROM `Enrollments`) `t`
WHERE `r` = 1