-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Dec 07, 2019 at 09:31 AM
-- Server version: 10.1.39-MariaDB
-- PHP Version: 7.3.5

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gogo`
--

-- --------------------------------------------------------

--
-- Table structure for table `privileges`
--

CREATE TABLE `privileges` (
  `id` int(11) NOT NULL,
  `role` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `privileges`
--

INSERT INTO `privileges` (`id`, `role`) VALUES
(1, 'admin'),
(2, 'petugas'),
(3, 'jamaah');

-- --------------------------------------------------------

--
-- Table structure for table `rekomendasi`
--

CREATE TABLE `rekomendasi` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `alamat` text NOT NULL,
  `foto` text NOT NULL,
  `lat` varchar(255) NOT NULL,
  `lng` varchar(255) NOT NULL,
  `rating` varchar(255) NOT NULL,
  `id_type` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `rekomendasi`
--

INSERT INTO `rekomendasi` (`id`, `nama`, `alamat`, `foto`, `lat`, `lng`, `rating`, `id_type`, `created_at`) VALUES
(1, 'test', 'alamat test', '098f6bcd4621d373cade4e832627b4f6\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191207005502.5r9DviyJR9tAiNTzMsaudiok.jpg', '-645557849', '342348', '4', 1, '0000-00-00 00:00:00'),
(2, 'test lagi', 'alamat test lagi', 'cdaf9b7f5cde730bd66b4f9715e79889\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20191207005631.0rGGgzPVCq4RDxPQfsaudiok.jpg', '-56539', '68946', '5', 1, '0000-00-00 00:00:00'),
(3, 'test lagi ya', 'alamat test lagi ya', 'f52b95ff7d2e0df2d0d7fc319b595444\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191207005711.7r9DviyJR9tAiNTzMsaudiok.jpg', '-56539', '68946', '5', 1, '2019-12-06 17:57:11'),
(4, 'test lagi ya yaya', 'alamat test lagi ya yaya', '5eaa8fef77d423a14bad15c8025f910a\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20191207005738.8rGGgzPVCq4RDxPQfsaudiok.jpg', '-7563', '68579253', '4', 2, '2019-12-06 17:57:38'),
(5, 'test lagi ya yaya terakhir', 'alamat test lagi ya yaya terakhir', '5d1bf4217a02eb55dde970429d8ec2e5\\e6rYbxNN9SBmUutoXgTQcvMvjxcEEF5eHBqpLmum20191207005800.5tmNEjWYbP21Y4ZGtsaudiok.jpg', '-7563', '68579253', '5', 2, '2019-12-06 17:58:00'),
(6, 'test lagi ya yaya terakhir kok', 'alamat test lagi ya yaya terakhir kok', 'fe5caab7d78eb3e6e0df0f03033b5956\\YUkQTt80Ce3SO7NGm0Qp64x0QMZCRgBFqTHbi5EY20191207005842.3tRW3x7liGOBWFFJtsaudiok.jpg', '-65986938', '5459.92', '5', 3, '2019-12-06 17:58:42');

-- --------------------------------------------------------

--
-- Table structure for table `typerekom`
--

CREATE TABLE `typerekom` (
  `id` int(11) NOT NULL,
  `type_rekom` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `typerekom`
--

INSERT INTO `typerekom` (`id`, `type_rekom`) VALUES
(1, 'Hotel & Ressort'),
(2, 'Halal Food'),
(3, 'Travel');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `tgl_lahir` date NOT NULL,
  `no_ktp` int(16) NOT NULL,
  `no_hp` int(15) NOT NULL,
  `no_visa` varchar(255) NOT NULL,
  `no_passpor` varchar(255) NOT NULL,
  `foto` varchar(255) NOT NULL,
  `id_privileges` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `nama`, `username`, `password`, `tgl_lahir`, `no_ktp`, `no_hp`, `no_visa`, `no_passpor`, `foto`, `id_privileges`, `created_at`) VALUES
(1, 'ica tester', 'test', '8XaAAOiU-OYnftQNMzEvPAgD1ghwNO9PYzpwfC8fyV8', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(2, 'ica1', 'ica1', '', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(3, 'ica12', 'ica12', '', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(4, 'ica123', 'ica123', '', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(5, 'ica1234', 'ica1234', '4aABSCgO6hPUrhTnUco4H9iBoV2yy-Brl_oDoB3hfgc', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(6, 'ica tester', 'test', 'MIr_eHkanblKmB6OnZ3nIIh4itV-Ibe_DpIXAgEaNo4', '0000-00-00', 0, 0, '', '', '', 0, '2019-12-06 14:58:17'),
(7, 'data lengkap', 'data', 'uP8XbdGZWumwFmnZRwQV88jtQsT5FGq-ieJ4la_W2yc', '1999-01-20', 2147483647, 887878529, '23957203572727298', '73647hhfuieyrui73', '8d777f385d3dfec8815d20f7496026dc\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191206224047.0r9DviyJR9tAiNTzMselamatpagi.jpg', 3, '2019-12-06 15:40:47'),
(8, 'data lengkap lagi ya', 'datalagi', 'PA8ge_kyOeoXgH8eO6omZL_iMHBMa78cHl-qrnwUXz4', '1993-04-10', 2147483647, 2147483647, '323732757584759748', 'er784577wgfw9', '557b3b8f78a619a7710b62b893fabfb4\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191206230705.7r9DviyJR9tAiNTzMselamatpagi.jpg', 3, '2019-12-06 16:07:05'),
(9, 'admin', 'admin', 'ijqz7PEJDUT3o0ZILnaWAutvGDvT49oA68DyLXIpsOY', '1993-04-10', 2147483647, 2147483647, '323732757584759748', 'er784577wgfw9', '21232f297a57a5a743894a0e4a801fc3\\XpdUC2sK8TpbNgR1osuS245jyPstpwGiLWZ9EZxB20191206231140.0r9DviyJR9tAiNTzMselamatpagi.jpg', 1, '2019-12-06 16:11:40'),
(10, 'petugas', 'petugas', 'pqNqVbY_0Sa_9roso6wDABNEEzRzOs-BFD6lEu0hmts', '1993-04-10', 2147483647, 2147483647, '4273657263924572387', '27678rebr', 'afb91ef692fd08c445e8cb1bab2ccf9c\\hKjM97yIOh3ftWEA68Fq8vhAbv0kzBOBH8Bj1LCv20191206231248.0rGGgzPVCq4RDxPQfselamatpagi.jpg', 2, '2019-12-06 16:12:48');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `privileges`
--
ALTER TABLE `privileges`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `rekomendasi`
--
ALTER TABLE `rekomendasi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `typerekom`
--
ALTER TABLE `typerekom`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `privileges`
--
ALTER TABLE `privileges`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `rekomendasi`
--
ALTER TABLE `rekomendasi`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `typerekom`
--
ALTER TABLE `typerekom`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
