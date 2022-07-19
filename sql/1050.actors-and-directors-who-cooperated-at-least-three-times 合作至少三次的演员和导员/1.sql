# Link: https://leetcode.cn/problems/actors-and-directors-who-cooperated-at-least-three-times

SELECT `actor_id`, `director_id`
FROM `actordirector`
GROUP BY `actor_id`, `director_id`
HAVING COUNT(*) >= 3