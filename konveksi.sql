-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 02, 2020 at 04:21 PM
-- Server version: 10.4.11-MariaDB
-- PHP Version: 7.4.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `konveksi`
--

-- --------------------------------------------------------

--
-- Table structure for table `buktibayar`
--

CREATE TABLE `buktibayar` (
  `idtransaksi` varchar(12) NOT NULL,
  `username` varchar(15) NOT NULL,
  `buktibayar` longblob DEFAULT NULL,
  `totalharga` int(11) NOT NULL,
  `status` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `produk`
--

CREATE TABLE `produk` (
  `idproduk` varchar(12) NOT NULL,
  `namaproduk` varchar(40) NOT NULL,
  `harga` int(11) NOT NULL,
  `deskripsi` text NOT NULL,
  `status` int(1) NOT NULL,
  `gambar` longblob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `produk`
--

INSERT INTO `produk` (`idproduk`, `namaproduk`, `harga`, `deskripsi`, `status`, `gambar`) VALUES
('200501194434', 'gembok', 4900, 'gembok', 1, ''),
('200501194507', 'tas kulit', 15500, 'tas', 1, ''),
('200501194552', 'masker', 4800, 'masker', 1, ''),
('200501194610', 'fila', 12000, 'fila', 1, ''),
('200501194638', 'lili', 9900, 'lili', 1, '');

-- --------------------------------------------------------

--
-- Table structure for table `trans`
--

CREATE TABLE `trans` (
  `idtransaksi` varchar(12) NOT NULL,
  `username` varchar(12) NOT NULL,
  `json_str` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `totalharga` int(11) NOT NULL,
  `status` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `trans`
--

INSERT INTO `trans` (`idtransaksi`, `username`, `json_str`, `totalharga`, `status`) VALUES
('200502181819', 'asd', '[{\"Idproduk\":\"200501194638\",\"Namaproduk\":\"lili\",\"Harga\":\"9900\",\"Jumlah\":\"2\",\"Keterangan\":\"2\"},{\"Idproduk\":\"200501194610\",\"Namaproduk\":\"fila\",\"Harga\":\"12000\",\"Jumlah\":\"3\",\"Keterangan\":\"3\"}]', 55800, 1);

-- --------------------------------------------------------

--
-- Table structure for table `transaksi`
--

CREATE TABLE `transaksi` (
  `idtransaksi` varchar(12) NOT NULL,
  `idproduk` varchar(12) NOT NULL,
  `username` varchar(20) NOT NULL,
  `jumlah` int(11) NOT NULL,
  `keterangan` text NOT NULL,
  `contoh` longblob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `username` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `nama` varchar(30) NOT NULL,
  `telepon` varchar(15) NOT NULL,
  `alamat` varchar(200) NOT NULL,
  `status` int(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`username`, `password`, `nama`, `telepon`, `alamat`, `status`) VALUES
('admin', '0cc175b9c0f1b6a831c399e269772661', 'Admin', '123', '123', 0),
('asd', '7815696ecbf1c96e6894b779456d330e', 'ASD', '123', '123asddsa', 1),
('qwe', '76d80224611fc919a5d54f0ff9fba446', 'qwe', '123', '123', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `buktibayar`
--
ALTER TABLE `buktibayar`
  ADD PRIMARY KEY (`idtransaksi`);

--
-- Indexes for table `produk`
--
ALTER TABLE `produk`
  ADD PRIMARY KEY (`idproduk`);

--
-- Indexes for table `trans`
--
ALTER TABLE `trans`
  ADD PRIMARY KEY (`idtransaksi`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`username`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
