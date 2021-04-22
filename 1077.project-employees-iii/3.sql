with `tmp` as (
    SELECT `project_id`, P.`employee_id`, `experience_years`
    FROM `project` AS P
             INNER JOIN Employee AS E ON P.employee_id = E.employee_id
)
select `project_id`, `employee_id`
from `tmp`
where (`project_id`, `experience_years`) IN
      (SELECT `project_id`, max(`experience_years`) FROM `tmp` GROUP BY `project_id`);