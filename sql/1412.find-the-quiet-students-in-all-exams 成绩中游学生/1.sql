# Link: https://leetcode.cn/problems/find-the-quiet-students-in-all-exams

SELECT `s`.`student_id`,
       `s`.`student_name`
FROM `student` `s`
WHERE `s`.`student_id` IN
      (
          SELECT `student_id`
          FROM (
                   SELECT `student_id`,
                          `score`,
                          `exam_id`,
                          MAX(`score`) OVER (PARTITION BY `exam_id`) `max_score`,
                          MIN(`score`) OVER (PARTITION BY `exam_id`) `min_score`
                   FROM `exam`
               ) `e`
          GROUP BY `student_id`
          HAVING SUM(IF(`min_score` < `e`.`score` AND `e`.`score` < `max_score`, 1, 0))
                     = COUNT(DISTINCT `exam_id`)
      )
ORDER BY `s`.`student_id`
