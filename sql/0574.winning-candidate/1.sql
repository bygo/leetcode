# Title: Winning Candidate
# Link: https://leetcode-cn.com/problems/winning-candidate

SELECT `Name`
FROM (SELECT `CandidateId` FROM `vote` GROUP BY `CandidateId` ORDER BY COUNT(`CandidateId`) DESC LIMIT 1) `t`
         JOIN `Candidate` ON `Candidate`.`id` = `CandidateId`