package service

import (
	"context"
	"mpc/internal/model"
	"mpc/internal/repository"
	"mpc/pkg/errors"
	"mpc/pkg/ethereum"
	"mpc/pkg/logger"
	"mpc/pkg/tss"
	"strings"

	"github.com/google/uuid"
)

type WalletService struct {
	walletRepo *repository.WalletRepository
	tssClient  *tss.TSS
	ethClient  *ethereum.EthClient
}

func NewWalletService(walletRepo *repository.WalletRepository, tssClient *tss.TSS, ethClient *ethereum.EthClient) *WalletService {
	return &WalletService{
		walletRepo: walletRepo,
		tssClient:  tssClient,
		ethClient:  ethClient,
	}
}

func (s *WalletService) CreateWallet(ctx context.Context, userID uuid.UUID) (model.Wallet, string, error) {
	// Create Ethereum wallet
	shareData, addressHex, err := s.tssClient.CreateWallet(ctx, userID.String())
	if err != nil {
		logger.Error("Service:CreateWallet", err)
		return model.Wallet{}, "", err
	}
	addressHex = strings.ToLower(addressHex)

	// Create wallet in repository
	wallet, err := s.walletRepo.CreateWallet(ctx, userID, addressHex, []byte(""), "Default")
	if err != nil {
		logger.Error("Service:CreateWallet", err)
		return model.Wallet{}, "", err
	}
	return wallet, shareData, nil
}

func (s *WalletService) GetWalletByUserID(ctx context.Context, userID uuid.UUID) (model.Wallet, error) {
	wallets, err := s.walletRepo.GetWalletsByUserID(ctx, userID)
	if err != nil {
		logger.Error("Service:GetWalletByUserID", err)
		return model.Wallet{}, err
	}
	if len(wallets) == 0 {
		return model.Wallet{}, errors.ErrWalletNotFound
	}
	return wallets[0], nil
}

func (s *WalletService) GetBalanceByAddress(
	ctx context.Context,
	request model.GetBalanceRequest,
	userID uuid.UUID,
) (model.GetBalanceResponse, error) {
	wallet, err := s.GetWalletByUserID(ctx, userID)
	if err != nil {
		logger.Error("Service:GetBalance", err)
		return model.GetBalanceResponse{}, err
	}
	address := strings.ToLower(wallet.Address)
	balance, err := s.ethClient.GetBalance(ctx, address)
	if err != nil {
		logger.Error("Service:GetBalance", err)
		return model.GetBalanceResponse{}, err
	}

	return model.GetBalanceResponse{
		Address: wallet.Address,
		Balance: balance,
	}, nil
}
