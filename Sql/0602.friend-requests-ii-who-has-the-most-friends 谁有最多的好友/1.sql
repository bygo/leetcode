# Link: https://leetcode.cn/problems/friend-requests-ii-who-has-the-most-friends

SELECT `id`, SUM(`n`) `num`
FROM (SELECT `accepter_id` `id`, COUNT(*) `n`
      FROM `request_accepted`
      GROUP BY `accepter_id`
      UNION ALL
      SELECT `requester_id` `id`, COUNT(*) `n`
      FROM `request_accepted`
      GROUP BY `requester_id`) `t`
GROUP BY `id`
ORDER BY `num` DESC
LIMIT 1;
