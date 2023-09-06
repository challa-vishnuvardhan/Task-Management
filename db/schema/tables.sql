Drop Table If Exists `task_details`
Create Table `task_details`(
    `task_id` bigint(20) unsigned Not Null AUTO_INCREMENT,
    `user_id` bigint(20) unsigned NOT NULL,
    `task_name` text NOT NULL,
    `start_date` bigint(20) unsigned,
    `due_date` bigint(20) unsigned,
    `description` TEXT DEFAULT Null,
    `priority` ENUM('low', 'medium', 'high') NOT NULL,
    `status` ENUM('to do', 'in progress', 'done') DEFAULT "to do",
    `is_active` tinyint(1) DEFAULT 1,   
    `created_by` varchar(320) NOT NULL, 
    `created_at` bigint(20) unsigned NOT NULL,  
    `modified_by` varchar(320) DEFAULT NULL,   
    `modified_at` bigint(20) unsigned DEFAULT NULL,  
    PRIMARY KEY (`task_id`), 
)

Drop Table If Exists `user_details`
Create Table `user_details`(
    `user_id` bigint(20) unsigned Not Null AUTO_INCREMENT,
    `first_name`  varchar(320) NOT NULL,
    `last_name`  varchar(320) NOT NULL,
    `email_id`  varchar(320)  NOT NUL,
    `is_active` tinyint(1) DEFAULT 1,   
    `created_by` varchar(320) NOT NULL, 
    `created_at` bigint(20) unsigned NOT NULL,  
    `modified_by` varchar(320) DEFAULT NULL,   
    `modified_at` bigint(20) unsigned DEFAULT NULL,  
    PRIMARY KEY (`user_id`), 
)
