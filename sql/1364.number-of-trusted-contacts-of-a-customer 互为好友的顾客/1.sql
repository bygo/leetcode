# Link: https://leetcode-cn.com/problems/number-of-trusted-contacts-of-a-customer

SELECT `invoice_id`,
       `cu1`.`customer_name`,
       `price`,
       COUNT(`co`.`contact_name`)   `contacts_cnt`,
       COUNT(`cu2`.`customer_name`) `trusted_contacts_cnt`
FROM `invoices`                `i`
         JOIN      `customers` `cu1` ON `i`.`user_id` = `cu1`.`customer_id`
         LEFT JOIN `contacts`  `co` ON `i`.`user_id` = `co`.`user_id`
         LEFT JOIN `customers` `cu2` ON `co`.`contact_email` = `cu2`.`email`
GROUP BY `i`.`invoice_id`
ORDER BY `i`.`invoice_id`