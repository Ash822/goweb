package service

import (
	. "github.com/ash822/goweb/entity"
	mocks "github.com/ash822/goweb/mocks/repository"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
	"testing"
)

var (
	id = "1"
	msg = Message{
		Id: id,
		Text: "ABBA",
	}

	invalidMsg = Message{
		Text: "",
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

func TestSvc_Create2(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	testMsgSvc := GetInstance(mr)
	_, err := testMsgSvc.Create(&invalidMsg)

	g.Expect(err).Should(HaveOccurred())
	g.Expect(err.Error()).Should(Equal("the message text is invalid or not found"))
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

func TestSvc_Update2(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	testMsgSvc := GetInstance(mr)
	_, err := testMsgSvc.Update("", &invalidMsg)

	g.Expect(err).Should(HaveOccurred())
	g.Expect(err.Error()).Should(Equal("the id provided is invalid"))
}

func TestSvc_Update3(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	testMsgSvc := GetInstance(mr)
	_, err := testMsgSvc.Update("1", &invalidMsg)

	g.Expect(err).Should(HaveOccurred())
	g.Expect(err.Error()).Should(Equal("the message text is invalid or not found"))
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

func TestSvc_Delete2(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	testMsgSvc := GetInstance(mr)
	err := testMsgSvc.Delete("")

	g.Expect(err).Should(HaveOccurred())
	g.Expect(err.Error()).Should(Equal("the id provided is invalid"))
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

func TestSvc_FindAll2(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	testMsgSvc := GetInstance(mr)
	_, err := testMsgSvc.FindById("")

	g.Expect(err).Should(HaveOccurred())
	g.Expect(err.Error()).Should(Equal("the id provided is invalid"))
}

func TestSvc_FindAll(t *testing.T) {
	g := NewGomegaWithT(t)
	ctrl := gomock.NewController(t)
	mr := mocks.NewMockMessageRepository(ctrl)

	mr.EXPECT().FindAll().Return([]Message{msg}, nil)

	testMsgSvc := GetInstance(mr)
	result, err := testMsgSvc.FindAll()

	g.Expect(err).NotTo(HaveOccurred())
	g.Expect(len(result)).Should(Equal(1))
}
