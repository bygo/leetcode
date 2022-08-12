# Link: https://leetcode.cn/problems/reformat-department-table

SELECT `id`,
       SUM(IF(`month` = 'Jan', `revenue`, NULL)) `jan_revenue`,
       SUM(IF(`month` = 'Feb', `revenue`, NULL)) `feb_revenue`,
       SUM(IF(`month` = 'Mar', `revenue`, NULL)) `mar_revenue`,
       SUM(IF(`month` = 'Apr', `revenue`, NULL)) `apr_revenue`,
       SUM(IF(`month` = 'May', `revenue`, NULL)) `may_revenue`,
       SUM(IF(`month` = 'Jun', `revenue`, NULL)) `jun_revenue`,
       SUM(IF(`month` = 'Jul', `revenue`, NULL)) `jul_revenue`,
       SUM(IF(`month` = 'Aug', `revenue`, NULL)) `aug_revenue`,
       SUM(IF(`month` = 'Sep', `revenue`, NULL)) `sep_revenue`,
       SUM(IF(`month` = 'Oct', `revenue`, NULL)) `oct_revenue`,
       SUM(IF(`month` = 'Nov', `revenue`, NULL)) `nov_revenue`,
       SUM(IF(`month` = 'Dec', `revenue`, NULL)) `dec_revenue`
FROM `department`
GROUP BY `id`
ORDER BY `id`;