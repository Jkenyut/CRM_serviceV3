-- +migrate Up
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `crm_service`
--

-- --------------------------------------------------------

--
-- Table structure for table `actors`
--

CREATE TABLE `actors` (
  `id` bigint UNSIGNED NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int UNSIGNED DEFAULT '2',
  `verified` enum('true','false') DEFAULT 'false',
  `active` enum('true','false') DEFAULT 'false',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `actors`
--

INSERT INTO `actors` (`id`, `username`, `password`, `role_id`, `verified`, `active`, `created_at`, `updated_at`) VALUES
(1, 'superadmin', '$2a$12$26EAzIyDQ5YZKE6pAeHKoeugjQ98lkXNPOEFjJikPmQReEMK4snoW', 1, 'true', 'true', '2023-06-02 04:41:18', '2023-06-04 05:18:58');
-- --------------------------------------------------------

--
-- Table structure for table `actor_role`
--

CREATE TABLE `actor_role` (
  `id` int UNSIGNED NOT NULL,
  `role_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `actor_role`
--

INSERT INTO `actor_role` (`id`, `role_name`) VALUES
(2, 'admin'),
(1, 'superadmin');

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `id` bigint UNSIGNED NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) NOT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

--
-- Table structure for table `register_approval`
--

CREATE TABLE `register_approval` (
  `id` bigint UNSIGNED NOT NULL,
  `admin_id` bigint UNSIGNED DEFAULT NULL,
  `super_admin_id` bigint UNSIGNED DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `register_approval`
--


-- Indexes for dumped tables
--

--
-- Indexes for table `actors`
--
ALTER TABLE `actors`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `actor_role` (`role_id`),
  ADD KEY `idx_username_actor` (`username`);

--
-- Indexes for table `actor_role`
--
ALTER TABLE `actor_role`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `role_name` (`role_name`);

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`),
  ADD KEY `idx_first_name` (`first_name`),
  ADD KEY `idx_last_name` (`last_name`),
  ADD KEY `customer_index_email` (`email`);

--
-- Indexes for table `register_approval`
--
ALTER TABLE `register_approval`
  ADD PRIMARY KEY (`id`),
  ADD KEY `register_admin_actor_role` (`admin_id`),
  ADD KEY `register_superadmin_actor_role` (`super_admin_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `actors`
--
ALTER TABLE `actors`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `actor_role`
--
ALTER TABLE `actor_role`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `customer`
--
ALTER TABLE `customer`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `register_approval`
--
ALTER TABLE `register_approval`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `actors`
--
ALTER TABLE `actors`
  ADD CONSTRAINT `actor_role` FOREIGN KEY (`role_id`) REFERENCES `actor_role` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

--
-- Constraints for table `register_approval`
--
ALTER TABLE `register_approval`
  ADD CONSTRAINT `register_admin_actor_role` FOREIGN KEY (`admin_id`) REFERENCES `actors` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `register_superadmin_actor_role` FOREIGN KEY (`super_admin_id`) REFERENCES `actors` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

-- +migrate Down
DROP TABLE IF EXISTS `actors`;