# Link: https://leetcode-cn.com/problems/daily-leads-and-partners

SELECT `date_id`,
       `make_name`,
       COUNT(DISTINCT `lead_id`)    `unique_leads`,
       COUNT(DISTINCT `partner_id`) `unique_partners`
FROM `dailysales`
GROUP BY `date_id`, `make_name`
