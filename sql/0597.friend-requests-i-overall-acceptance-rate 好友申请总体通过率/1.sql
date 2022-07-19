# Link: https://leetcode.cn/problems/friend-requests-i-overall-acceptance-rate

SELECT ROUND(
               IFNULL(
                           (SELECT COUNT(*)
                            FROM (SELECT DISTINCT `requester_id`, `accepter_id` FROM `requestaccepted`) AS `a`)
                           /
                           (SELECT COUNT(*)
                            FROM (SELECT DISTINCT `sender_id`, `send_to_id` FROM `friendrequest`) AS `b`),
                           0)
           , 2) `accept_rate`;