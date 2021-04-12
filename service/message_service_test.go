package service

import (
	"github.com/ash822/goweb/entity"
	mocks "github.com/ash822/goweb/mocks/repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"testing"
)

var (
	id = "1"
	msg = entity.Message{
		Id: id,
		Text: "ABBA",
	}
)

func TestSvc_Create(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().Create(&msg).Return(&msg, nil)

	testMsgSvc := GetInstance(mr)
	result, err := testMsgSvc.Create(&msg)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(result.Text).Should(Equal("ABBA"))
	g.Expect(result.Palindrome).To(BeTrue())
}

func TestSvc_Update(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().Update(&msg).Return(&msg, nil)

	testMsgSvc := GetInstance(mr)
	result, err := testMsgSvc.Update(id, &msg)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(result.Id).Should(Equal(id))
	g.Expect(result.Palindrome).To(BeTrue())
}

func TestSvc_Delete(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().Delete(id).Return(nil)

	testMsgSvc := GetInstance(mr)
	err := testMsgSvc.Delete(id)

	g.Expect(err).NotTo(HaveOccurred())
}

func TestSvc_FindById(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().FindById(id).Return(&msg, nil)

	testMsgSvc := GetInstance(mr)
	result, err := testMsgSvc.FindById(id)

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(result.Id).Should(Equal(id))
}

func TestSvc_FindAll(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().FindAll().Return([]entity.Message{msg}, nil)

	testMsgSvc := GetInstance(mr)
	result, err := testMsgSvc.FindAll()

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(result)).Should(Equal(1))
}
