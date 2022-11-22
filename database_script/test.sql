-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 22, 2022 at 07:07 AM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 7.4.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test`
--

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `timestamp` varchar(255) NOT NULL,
  `city` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `amount`, `timestamp`, `city`, `created_at`) VALUES
(45, 12000, '2022-11-18T12:4:51.312Z', '', '2022-11-18 06:34:59'),
(46, 12000, '2022-11-18T21:20:51.312Z', '', '2022-11-18 21:21:38'),
(47, 22, '', '', '2022-11-22 05:27:54'),
(48, 12000, '2022-11-22T05:39:51.312Z', '', '2022-11-22 05:40:13'),
(49, 12000, '2022-11-22T05:43:51.312Z', '', '2022-11-22 05:44:30'),
(50, 12000, '2022-11-22T05:45:51.312Z', '', '2022-11-22 05:46:04'),
(51, 1000, '2022-11-22T05:45:51.312Z', '', '2022-11-22 05:46:11'),
(52, 1000, '2022-11-22T05:45:51.312Z', '', '2022-11-22 05:46:13'),
(53, 10, '2022-11-22T05:45:51.312Z', '', '2022-11-22 05:46:18'),
(54, 10, '2022-11-22T05:51:51.312Z', 'Moradabad4', '2022-11-22 05:52:27'),
(55, 10, '2022-11-22T05:52:51.312Z', '', '2022-11-22 05:53:02'),
(56, 11, '2022-11-22T05:52:51.312Z', '', '2022-11-22 05:53:09'),
(57, 11, '2022-11-22T05:52:51.312Z', '', '2022-11-22 05:53:11');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=58;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
