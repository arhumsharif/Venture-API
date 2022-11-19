-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 19, 2022 at 02:31 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `venture`
--

-- --------------------------------------------------------

--
-- Table structure for table `user_details`
--

CREATE TABLE `user_details` (
  `id` int(10) NOT NULL,
  `user_guid` text NOT NULL,
  `email` text NOT NULL,
  `password` text NOT NULL,
  `secret_key` text NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp(),
  `is_deleted` int(2) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_details`
--

INSERT INTO `user_details` (`id`, `user_guid`, `email`, `password`, `secret_key`, `date_created`, `is_deleted`) VALUES
(1, '3213232323', 'arhum', '1234', '', '2022-11-16 22:05:45', 0),
(2, '3213232323', 'talha', '0123', '', '2022-11-17 23:07:33', 0),
(3, '3213232323', 'talha', '0123', '', '2022-11-17 23:10:26', 0),
(4, '3213232323', 'talha', '0123', '', '2022-11-17 23:19:23', 0),
(5, '3213232323', 'talha', '0123', '', '2022-11-17 23:20:12', 0),
(6, '1122', 'amjad', '1122', '', '2022-11-17 23:20:51', 0),
(7, '1122', 'amjad', '1122', '', '2022-11-17 23:27:15', 0),
(8, '111999', 'arhumsharif06', '0852', '', '2022-11-19 13:08:33', 0),
(9, '111999', 'arhumsharif06@gmail.com', '0852', '', '2022-11-19 13:09:04', 0),
(10, '222444', 'talha@gmail.com', '5656', '', '2022-11-19 13:11:19', 0),
(11, '222444', 'hassan@gmail.com', '6969', '', '2022-11-19 14:15:42', 0),
(12, '72d3dc11-103c-4c34-88a7-e9cb21a35647', 'noman@gmail.com', '6969', '', '2022-11-19 14:29:06', 0),
(13, '7b170e11-5337-4834-9773-f00924430384', 'noman@gmail.com', '6969', '', '2022-11-19 14:30:15', 0),
(14, '563b883b-f377-4f3e-8775-d6f4fa1f1124', 'babar@gmail.com', '56', '', '2022-11-19 14:34:43', 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_education`
--

CREATE TABLE `user_education` (
  `id` int(10) NOT NULL,
  `education_guid` text NOT NULL,
  `user_guid` text NOT NULL,
  `school` text NOT NULL,
  `user_from` text NOT NULL,
  `user_to` text NOT NULL,
  `degree` text NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp(),
  `is_deleted` int(2) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_education`
--

INSERT INTO `user_education` (`id`, `education_guid`, `user_guid`, `school`, `user_from`, `user_to`, `degree`, `date_created`, `is_deleted`) VALUES
(1, '7774416b-a4dd-4ed6-8a92-9968db033499', '1324357575', 'lasalle', '2022-09-21', '2022-10-21', 'bscs', '2022-11-19 15:11:01', 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_experience`
--

CREATE TABLE `user_experience` (
  `id` int(10) NOT NULL,
  `experience_guid` text NOT NULL,
  `user_guid` text NOT NULL,
  `job_title` text NOT NULL,
  `job_description` text NOT NULL,
  `company_name` text NOT NULL,
  `user_from` text NOT NULL,
  `user_to` text NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp(),
  `is_deleted` int(2) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_experience`
--

INSERT INTO `user_experience` (`id`, `experience_guid`, `user_guid`, `job_title`, `job_description`, `company_name`, `user_from`, `user_to`, `date_created`, `is_deleted`) VALUES
(1, '34b636bf-feaf-4873-8082-4cc085738b8d', '122454655', 'developer', 'a good developer', 'weteck', '2022-09-15', '2021-03-21', '2022-11-19 15:53:37', 0);

-- --------------------------------------------------------

--
-- Table structure for table `user_projects`
--

CREATE TABLE `user_projects` (
  `id` int(10) NOT NULL,
  `project_guid` text NOT NULL,
  `user_guid` text NOT NULL,
  `title` text NOT NULL,
  `description` text NOT NULL,
  `technologies` text NOT NULL,
  `date_created` datetime NOT NULL DEFAULT current_timestamp(),
  `is_deleted` int(2) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user_projects`
--

INSERT INTO `user_projects` (`id`, `project_guid`, `user_guid`, `title`, `description`, `technologies`, `date_created`, `is_deleted`) VALUES
(1, 'f6c7023c-8257-4723-8084-8a84c0db7987', '2143545335', 'FYP', 'my desc', 'c++', '2022-11-19 16:06:40', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user_details`
--
ALTER TABLE `user_details`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_education`
--
ALTER TABLE `user_education`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_experience`
--
ALTER TABLE `user_experience`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_projects`
--
ALTER TABLE `user_projects`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `user_details`
--
ALTER TABLE `user_details`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `user_education`
--
ALTER TABLE `user_education`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user_experience`
--
ALTER TABLE `user_experience`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `user_projects`
--
ALTER TABLE `user_projects`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
