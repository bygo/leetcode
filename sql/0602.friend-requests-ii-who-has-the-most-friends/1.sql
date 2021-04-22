# Title: Friend Requests II: Who Has the Most Friends
# Link: https://leetcode-cn.com/problems/friend-requests-ii-who-has-the-most-friends

SELECT `id`, SUM(`n`) AS `num`
FROM (SELECT `accepter_id` AS `id`, COUNT(*) AS `n`
      FROM `request_accepted`
      GROUP BY `accepter_id`
      UNION ALL
      SELECT `requester_id` AS `id`, COUNT(*) AS `n`
      FROM `request_accepted`
      GROUP BY `requester_id`) `t`
GROUP BY `id`
ORDER BY `num` DESC
LIMIT 1;
