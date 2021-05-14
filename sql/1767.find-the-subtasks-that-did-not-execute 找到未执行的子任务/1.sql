# Link: https://leetcode-cn.com/problems/find-the-subtasks-that-did-not-execute

WITH RECURSIVE `t`(`task_id`, `subtask_id`) AS (
    SELECT `task_id`, `subtasks_count`
    FROM `tasks`
    UNION ALL
    SELECT `task_id`, `subtask_id` - 1
    FROM `t`
    WHERE `subtask_id` - 1 > 0
)

SELECT *
FROM `t`
         LEFT JOIN `executed` USING (`task_id`, `subtask_id`)
WHERE `executed`.`subtask_id` IS NULL
ORDER BY `task_id`, `subtask_id`