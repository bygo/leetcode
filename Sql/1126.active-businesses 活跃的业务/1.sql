# Link: https://leetcode-cn.com/problems/active-businesses

SELECT `business_id`
FROM (SELECT *,
             AVG(`occurences`) OVER (PARTITION BY `event_type`) `a`
      FROM `events`) `t`
WHERE `occurences` > `a`
GROUP BY `business_id`
HAVING COUNT(*) > 1