-- Query 1: Direct profile count
SELECT COUNT(*) as profile_count
FROM dcdd_user_data.dcdd_user_profiles
WHERE dzongkhag_id = '9530433c-f377-4af1-8097-08b0684180fe'
  AND eccd_id = '007a3db7-0499-4900-aeed-5eaf948de943';

-- Query 2: Users with those profiles  
SELECT COUNT(DISTINCT u.id) as user_count
FROM dcdd_auth.dcdd_users u
JOIN dcdd_user_data.dcdd_user_profiles up ON up.user_id = u.id
WHERE up.dzongkhag_id = '9530433c-f377-4af1-8097-08b0684180fe'
  AND up.eccd_id = '007a3db7-0499-4900-aeed-5eaf948de943';

-- Query 3: Users with activity records in that ECCD
SELECT COUNT(DISTINCT u.id) as user_count_with_activity
FROM dcdd_auth.dcdd_users u
JOIN dcdd_user_data.dcdd_user_profiles up ON up.user_id = u.id
WHERE up.dzongkhag_id = '9530433c-f377-4af1-8097-08b0684180fe'
  AND up.eccd_id = '007a3db7-0499-4900-aeed-5eaf948de943'
  AND u.id IN (SELECT user_id FROM dcdd_auth.dcdd_user_activities);
