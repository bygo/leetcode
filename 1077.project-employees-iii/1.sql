
//Title: Project Employees III
//Link: https://leetcode-cn.com/problems/project-employees-iii

# Write your MySQL query statement below

SELECT
	p.project_id,
	p.employee_id
FROM
	Project AS p
	INNER JOIN Employee AS e
		ON p.employee_id = e.employee_id
WHERE (p.project_id, e.experience_years) IN (
	SELECT
		p.project_id,
		MAX(e.experience_years)
	FROM
		Project AS p
		INNER JOIN Employee AS e
			ON p.employee_id = e.employee_id
	GROUP BY p.project_id
);


SELECT PROJECT_ID, EMPLOYEE_ID
FROM (SELECT PROJECT_ID, P.EMPLOYEE_ID, RANK() OVER(PARTITION BY PROJECT_ID ORDER BY EXPERIENCE_YEARS DESC) AS `RANK`
      FROM PROJECT AS P
      INNER JOIN EMPLOYEE AS E
      ON P.EMPLOYEE_ID = E.EMPLOYEE_ID) AS A
WHERE `RANK` = 1;


with `tmp` as (
SELECT * FROM `project` AS p INNER JOIN Employee AS E ON p.employee_id = e.employee_id
)
select `project_id`,`employee_id` from `tmp` where `experience_years` = (SELECT max(`experience_years`) FROM `tmp`);