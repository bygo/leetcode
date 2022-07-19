# Link: https://leetcode.cn/problems/highest-grade-for-each-student

SELECT `student_id`, `course_id`, `grade`
FROM (SELECT *, ROW_NUMBER() OVER (PARTITION BY `student_id` ORDER BY `grade` DESC,`course_id`) `r`
      FROM `enrollments`) `t`
WHERE `r` = 1